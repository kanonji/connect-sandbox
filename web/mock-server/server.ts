import { ConnectRouter } from "@bufbuild/connect"
import { connectNodeAdapter } from "@bufbuild/connect-node"
import { ElizaService } from "../gen/eliza/v1/eliza_connect"
import {
  SayRequest,
  IntroduceRequest,
  ConverseRequest,
} from "../gen/eliza/v1/eliza_pb"
import { Empty } from "@bufbuild/protobuf"
import http2 from "http2"
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
    res.writeHead(404)
    res.end()
  },
})

http2
  .createServer(handler)
  .listen(3000, () => {
    stdout.write("The server is listening on http://localhost:3000\n")
    stdout.write("Run `npm run client` for a terminal client.\n")
  })
