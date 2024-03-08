
import {ChannelCredentials} from "@grpc/grpc-js";
import {GrpcTransport} from "@protobuf-ts/grpc-transport";
import {UserServiceClient} from "./gen/js/user/v1/user.client";

let transport = new GrpcTransport({
    host: "localhost:50055",
    channelCredentials: ChannelCredentials.createInsecure(),
});

let client = new UserServiceClient(transport);

const getResponse = async () => {
    let {response} = await client.passwordLogin({
        email: "test@example.com",
        password: "password"
    });
    console.log("got a response! ", response)

}

getResponse().catch(e => console.error(e));
