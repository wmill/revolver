const Hapi = require('hapi');
const HapiJwt = require('hapi-auth-jwt2');

const server = Hapi.server({
  port: 3000,
  // Other server configurations...
});

// Register the hapi-auth-jwt2 plugin
await server.register(HapiJwt);

// Define the JWT authentication strategy
server.auth.strategy('jwt', 'jwt', {
  key: 'yourSecretKey', // Change this to your own secret key
  validate: async (decoded, request) => {
    // Custom validation logic to verify the token
    // You can perform database lookups, etc., and return an object containing the user information
    // For example: { isValid: true, credentials: { id: decoded.userId } }
  },
  verifyOptions: {
    algorithms: ['HS256'], // Specify the JWT signing algorithm you are using
  },
});

// Set the default authentication strategy to 'jwt'
server.auth.default('jwt');

export {}