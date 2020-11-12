const fetching = {
    'GET /fetching/fetch': 'FetchingController.fetch',
    'GET /fetching/aggregate': 'FetchingController.aggregate',
    'GET /fetching/validate-token': 'FetchingController.validateToken',
}

const routes = {
    ...fetching,
} 

module.exports = routes;
