class Comodity {
    constructor(json, usd) {
        this.UUID = json.UUID;
        this.komoditas = json.komoditas;
        this.area_provinsi = json.area_provinsi;
        this.area_kota = json.area_kota;
        this.size = json.size;
        this.price = json.price;
        this.price_usd = usd;
        this.tgl_parsed = json.tgl_parsed;
        this.timestamp = json.timestamp;
    }
}

module.exports = Comodity;