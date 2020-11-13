const fetching = {
    'GET /fetching/fetch': 'FetchingController.fetchAll',
    'GET /fetching/aggregate': 'FetchingController.aggregate',
    'GET /fetching/validate-token': 'AuthController.validateToken',
}

const routes = {
    ...fetching,
} 

module.exports = routes;
