**Infuse** is the gateway between your app and hosted LLMs. It streamlines API requests to dozens of hosted LLMs with a standarized RESTful API. <br><br>

✅&nbsp; **Load balance** across multiple models, providers, and keys <br> 
✅&nbsp; **Fallbacks** make sure your app stays resilient  <br>
✅&nbsp; **Automatic Retries** with exponential fallbacks come by default  <br>
✅&nbsp; Plug-in middleware as needed <br>
<br>

## Supported Providers [WIP]

| Provider  | Support | Stream | Supported Endpoints |
| OpenAI | ✅  |✅  | `/completions`, `/chat/completions`,`/embeddings` |
| Azure OpenAI | ✅  |✅  | `/completions`, `/chat/completions`,`/embeddings` |
| Anyscale | ✅   | ✅  | `/chat/completions` |
| Google Gemini & Palm | ✅  |✅  | `/generateMessage`, `/generateText`, `/embedText` |
| Anthropic  | ✅  |✅  | `/messages`, `/complete` |
| Cohere  | ✅  |✅  | `/generate`, `/embed` |
| Together AI  | ✅  |✅  | `/chat/completions`, `/completions`, `/inference` |
| Perplexity  | ✅  |✅  | `/chat/completions` |
| Mistral  | ✅  |✅  | `/chat/completions`, `/embeddings` |
