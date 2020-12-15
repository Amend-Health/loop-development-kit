using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using Google.Protobuf;
using Google.Protobuf.Collections;
using Grpc.Core;
using OliveHelpsLDK.Logging;
using Proto;

namespace OliveHelpsLDK.Network
{
    internal class NetworkClient : BaseClient<Proto.Network.NetworkClient>, INetworkService
    {
        internal NetworkClient(Proto.Network.NetworkClient client, Session session, ILogger logger) : base(client,
            session, logger, "network")
        {
        }

        internal NetworkClient(ChannelBase channelBase, Session session, ILogger logger) : this(
            new Proto.Network.NetworkClient(channelBase), session, logger)
        {
        }


        public Task<HTTPResponse> HTTPRequest(HTTPRequest request, CancellationToken cancellationToken = default)
        {
            var message = ToProto(request);

            return Client.HTTPRequestAsync(message, CreateOptions(cancellationToken)).ResponseAsync
                .ContinueWith(LoggedParser<Task<HTTPResponseMsg>, HTTPResponse>(task => FromProto(task.Result)),
                    cancellationToken);
        }

        private HTTPRequestMsg ToProto(HTTPRequest request)
        {
            var message = new HTTPRequestMsg
            {
                Session = CreateSession(),
                Body = ByteString.CopyFrom(request.Body),
                Method = request.Method,
                Url = request.URL
            };        
            
            AddHeadersToMessage(message, request.Headers);
            
            return message;
        }
        
        private static HTTPRequestMsg AddHeadersToMessage(HTTPRequestMsg message, IDictionary<string, IList<string>> requestHeaders)
        {
            foreach (var kvp in requestHeaders)
            {
                var header = new Header();
                header.Values.Add(kvp.Value);
                message.Headers.Add(kvp.Key, header);
            }

            return message;
        }

        internal static HTTPResponse FromProto(HTTPResponseMsg response)
        {
            var clientResponse = new HTTPResponse
            {
                Data = response.ToByteArray(),
                ResponseCode = checked((int) response.ResponseCode),
                Headers = ParseHeaders(response.Headers)
            };

            return clientResponse;
        }

        private static IDictionary<string, IList<string>> ParseHeaders(MapField<string, Header> headers)
        {
            var parsedHeaders = new Dictionary<string, IList<string>>();

            foreach (var kvp in headers)
            {
                var values = new List<string>(kvp.Value.Values);
                parsedHeaders.Add(kvp.Key, values);
            }

            return parsedHeaders;
        }
    }
}