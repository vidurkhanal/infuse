[Infuse] is the gateway between your app and hosted LLMs. It streamlines API requests to dozens of hosted LLMs with a standarized RESTful API. 

✅&nbsp; **Load balance** across multiple models, providers, and keys <br> 
✅&nbsp; **Fallbacks** make sure your app stays resilient  <br>
✅&nbsp; **Automatic Retries** with exponential fallbacks come by default  <br>
✅&nbsp; Plug-in middleware as needed <br>
<br>

## Supported Providers

|| Provider  | Support | Stream | Supported Endpoints |
|---|---|---|---|--|
| <img src="docs/images/openai.png" width=25 />| OpenAI | ✅  |✅  | `/completions`, `/chat/completions`,`/embeddings` |
| <img src="docs/images/azure.png" width=25>| Azure OpenAI | ✅  |✅  | `/completions`, `/chat/completions`,`/embeddings` |
| <img src="docs/images/anyscale.png" width=25>| Anyscale | ✅   | ✅  | `/chat/completions` |
| <img src="https://upload.wikimedia.org/wikipedia/commons/2/2d/Google-favicon-2015.png" width=25>| Google Gemini & Palm | ✅  |✅  | `/generateMessage`, `/generateText`, `/embedText` |
| <img src="docs/images/anthropic.png" width=25>| Anthropic  | ✅  |✅  | `/messages`, `/complete` |
| <img src="docs/images/cohere.png" width=25>| Cohere  | ✅  |✅  | `/generate`, `/embed`, `/rerank` |
| <img src="https://assets-global.website-files.com/64f6f2c0e3f4c5a91c1e823a/654693d569494912cfc0c0d4_favicon.svg" width=25>| Together AI  | ✅  |✅  | `/chat/completions`, `/completions`, `/inference` |
| <img src="https://www.perplexity.ai/favicon.svg" width=25>| Perplexity  | ✅  |✅  | `/chat/completions` |
| <img src="https://docs.mistral.ai/img/favicon.ico" width=25>| Mistral  | ✅  |✅  | `/chat/completions`, `/embeddings` |
