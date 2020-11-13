const bodyParser = require('body-parser');
const express = require('express');
const helmet = require('helmet');
const http = require('http');
const mapRoutes = require('express-routes-mapper');
const cors = require('cors');

const config = require('./fetching-api/config/');
const auth = require('./fetching-api/policies/auth.policy')

const app = express();
const server = http.Server(app);
const mappedRoutes = mapRoutes(config.routes, 'fetching-api/controllers/');

app.use(cors());

app.use(helmet({
    dnsPrefetchControl: false,
    frameguard: false,
    ieNoOpen: false,
}))

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.all('/*', (req, res, next) => auth(req, res, next));

app.use('/', mappedRoutes);

server.listen(config.port);