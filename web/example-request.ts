
import {ChannelCredentials} from "@grpc/grpc-js";
import {GrpcTransport} from "@protobuf-ts/grpc-transport";
// import {ExampleServiceClient, IExampleServiceClient} from "./service-example.client";
// import {ExampleRequest, FailRequest} from "./service-example";

import {UserServiceClient} from "./gen/js/user/v1/user.client";

let transport = new TwirpFetchTransport({
    baseUrl: "localhost:4000"
});

let client = new HaberdasherClient(transport);

let {response} = await client.makeHat({ inches: 11 });
console.log("got a small hat! " + response)