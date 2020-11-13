const jwt = require("jsonwebtoken");
const secret = process.env.TOKEN_SECRET;

const AuthController = () => {
    const validateToken = async (req,res) => {
        const token = req.headers.authorization.split(' ')[1];
        try {
            const claims = jwt.verify(token, secret);
            return res.status(201).json({ claims });
        } catch (err) {
            return res.status(400);
        }
    }
    return {
        validateToken,
    }
}

module.exports = AuthController;