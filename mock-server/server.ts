import { ConnectRouter } from "@bufbuild/connect"
import { connectNodeAdapter } from "@bufbuild/connect-node"
import { ElizaService } from "../web/gen/eliza/v1/eliza_connect"
import {
  SayRequest,
  IntroduceRequest,
  ConverseRequest,
} from "../web/gen/eliza/v1/eliza_pb"
import { Empty } from "@bufbuild/protobuf"
import http2 from "http2"
import { readFileSync } from "fs";
import { stdout } from "process"

// Let's implement our RPCs and add them to the Connect router:
function routes(router: ConnectRouter) {
  router.service(ElizaService, {
    say(req: SayRequest) {
      return {
        sentence: `You said ${req.sentence}`,
      }
    },
    async *introduce(req: IntroduceRequest) {
      yield { sentence: `Hi ${req.name}, I'm eliza` }
      await delay(250)
      yield {
        sentence: `Before we begin, ${req.name}, let me tell you something about myself.`,
      }
      await delay(150)
      yield { sentence: `I'm a Rogerian psychotherapist.` }
      await delay(150)
      yield { sentence: `How are you feeling today?` }
    },
    async *converse(reqs: AsyncIterable<ConverseRequest>) {
      for await (const req of reqs) {
        yield {
          sentence: `You said ${req.sentence}`,
        }
      }
    },
  })
  function delay(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms))
  }
}

// The adapter turns our RPC routes into as Node.js request handler.
const handler = connectNodeAdapter({
  routes,
  // If none of the RPC routes match, this handler is called.
  // We serve our web interface here:
  fallback(req, res) {
    switch (req.url) {
      case "/":
        res.writeHead(200, { "content-type": "text/html" })
        res.write("200 OK", "utf8")
        res.end();
        break;
      default:
        res.writeHead(404);
        res.write("404 Not found", "utf8")
        res.end();
    }
  },
})

http2
  .createSecureServer(
    {
      // We configure the server to use the locally-trusted development certificate
      // we have created with mkcert:
      key: readFileSync("localhost+2-key.pem", "utf8"),
      cert: readFileSync("localhost+2.pem", "utf8"),
      // Because we are using a certificate, we can use ALPN to offer both HTTP 1.1
      // and HTTP/2 on the same port.
      allowHTTP1: true,
    },
    handler
  )
  .listen(3000, () => {
    stdout.write("The server is listening on https://localhost:3000\n")
  })
