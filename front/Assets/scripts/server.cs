using System;
using Grpc.Core;
using GrpcBuffer;

namespace csharp.client
{
    internal class Program
    {
        static private gameServe Client;
        private Program() {
            channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
            Client = new gameServe.gameServeClient(channel);
            channel.ShutdownAsync().Wait();
        }

        public gameServe getClient() {
            return Client;
        }
    }
}