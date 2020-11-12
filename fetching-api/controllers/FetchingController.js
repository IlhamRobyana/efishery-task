const { TokenExpiredError } = require("jsonwebtoken");

const FetchingController = () => {
    const fetch = async (req,res) => {
        const { body } = req;
        return
    };
    const aggregate = async (req,res) => {
        return
    };
    const validateToken = async (req,res) => {
        return
    }
    return {
        fetch,
        aggregate,
        validateToken,
    }
}

module.exports = FetchingController;