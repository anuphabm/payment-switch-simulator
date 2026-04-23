import tcp from 'k6/x/tcp';
import { sleep } from 'k6';

export const options = {
    vus: 20,          // concurrent users
    duration: '30s',  // test time
};

export default function () {
    const socket = new tcp.Socket()

    socket.on("connect", () => {
        console.log("Connected")

        const traceId = Math.floor(Math.random() * 1000000);
        const payload = JSON.stringify({
            mti: "0200",
            trace_id: traceId.toString(),
            amount: 100
        }) + "\n";

        socket.write(payload);
    })

    socket.on("data", (data) => {
        console.log("Received data")
        const str = String.fromCharCode.apply(null, new Uint8Array(data))
        console.log(str)
        socket.destroy()
    })


    // const res = conn.read();

    // // optional check
    // if (!res.includes('"response_code":"00"') &&
    //     !res.includes('"response_code":"91"')) {
    //     console.error("unexpected response:", res);
    // }

    // conn.close();

    sleep(0.1);
}