const authService = require('../services/auth.service');
const currencyService = require("../services/currency.service");
const Comodity = require('../models/Comodity');
const fetch = require('node-fetch');
const url = 'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list';
const settings = { method: "Get" };
const jsonAggregate = require('json-aggregate')
const urlCurrency = 'https://free.currconv.com/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=b3eb57a74dcf08edf175'
const redis = require("redis");
const mathjs = require("mathjs")

let redisClient = redis.createClient(process.env.REDIS_PORT, process.env.REDIS_HOST);


const FetchingController = () => {
    const fetchAll = async (req,res) => {
        try {
            var usd_rate = 0;
            redisClient.get("idr_usd", function(err, val) {
                if (val != null) {
                    usd_rate = parseFloat(val);
                } else {
                    try {
                        fetch(urlCurrency, settings)
                        .then(res => res.json())
                        .then((json) => {
                            redisClient.set("idr_usd", json.IDR_USD);
                            console.log(json.IDR_USD);
                            usd_rate = parseFloat(json.IDR_USD);
                        });
                    } catch (err) {
    
                    }
                }
            });
            // I didn't make the currency conversion a method because it always
            // returns an empty value
            fetch(url, settings)
                .then(res => res.json())
                .then((json) => {
                    var response = [];
                    json.forEach(obj => {
                        console.log(obj.price)
                        console.log(usd_rate)
                        let usd = toString(obj.price * usd_rate)
                        let comodity = new Comodity(obj, usd)
                        response.push(comodity)
                    })
                    return res.status(200).json({ response });
                });
        } catch (err) {
            return res.status(400);
        }
    };
    const aggregate = async (req,res) => {
        if (await authService().verifyHasRoleAdmin(req)) {
            try {
                fetch(url, settings)
                    .then(res => res.json())
                    .then((json) => {
                        hash = json.reduce((p, c) => (p[c.area_provinsi] ? p[c.area_provinsi].push(c) : p[c.area_provinsi] = [c],p), {}),
                        grouped = Object.keys(hash).map(k => ({area_provinsi: k, komoditas: hash[k]}))
                        grouped.forEach(obj => {
                            var list = [];
                            obj.komoditas.forEach( obj2 => {
                                list.push(parseInt(obj2.price))
                                delete obj2
                            })
                            obj.min = mathjs.min(list)
                            obj.max = mathjs.max(list)
                            obj.avg = mathjs.mean(list)
                            obj.median = mathjs.median(list)
                        })
                        // var stringified = JSON.stringify(json)
                        // const collection = jsonAggregate.create(stringified);
                        // var grouped = collection.group({
                        //     id : ['area_provinsi', 'tgl_parsed'],
                        //     min : { $min: 'price' },
                        //     max : { $max: 'price'},
                        //     average : { $avg: 'price'}
                        // }).exec()
                        return res.status(200).json({ grouped });
                    });
            } catch (err) {
                return res.status(400);
            }
        } else {
            return res.status(401).json({ msg: 'Unauthorized' });
        }
    };
    

    return {
        fetchAll,
        aggregate,
    }
}

module.exports = FetchingController;