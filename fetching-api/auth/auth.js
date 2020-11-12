const jwt = require('jsonwebtoken')
const secret = process.env.JWT_SECRET;

const auth = () => {
    const verifyToken = (token, cb) => jwt.verify(token, secret, {}, cb);
    const verifyHasRoleAdmin = async (req) => {
        const token = req.headers.authorization.split(' ')[1];
        try {
            const decoded = jwt.verify(token, secret)
            if (decoded.role === 'ADMIN') {
                return true;
            }
            return false;
        } catch (err) {
            return false;
        }
    }
    return {
        verifyToken,
        verifyHasRoleAdmin,
    };
}

module.exports = auth;