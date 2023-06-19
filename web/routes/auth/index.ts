import {Plugin} from "@hapi/hapi";

// function register (server, options) {
//     server.route([{
//         method: "POST",
//         path: "/login",
//         config: { auth: false },
//         handler: function (request, h) {
//             return { text: "Token not required" };
//         },
//     }])
//     return Promise.resolve();
// };
  
// register.attributes = {name: "authroutes", version: "" }

const plugin = Plugin<object> : {
    register: (server, options) => {
        server.route([{
            method: "POST",
            path: "/login",
            config: { auth: false },
            handler: function (request, h) {
                return { text: "Token not required" };
            },
        }])
        return Promise.resolve();
    } 
}

export default plugin;