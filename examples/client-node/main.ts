/**
 * Example usage of wechat-mem0-core gRPC client
 *
 * First start the server:
 *   cd wechat-mem0-core && ./bin/wechat-mem0-core --addr=:50051
 *
 * Then run this example:
 *   cd client-node && npm install && node examples/example.mjs
 */

import GRPCClient from './client';

async function main() {
    const client = new GRPCClient('localhost:50051');

    try {
        // Connect to the gRPC server
        console.log('Connecting to gRPC server...');
        client.connect();
        console.log('Connected!');

        // Example: Set log level
        console.log('Setting log level to debug...');
        await client.setLogLevel('debug');
        console.log('Log level set successfully');

        // Example: Get WeChat instances
        console.log('Getting WeChat instances...');
        const response = await client.getWeChatInstances();
        console.log('WeChat instances:', response.accounts);

        // Example: Run with config (will return "not implemented" from stub)
        // console.log('Running with config...');
        // await client.run('/path/to/config.yaml');

    } catch (err) {
        console.error('Error:', err.message);
    } finally {
        // Close the connection
        client.close();
        console.log('Connection closed');
    }
}

main().catch(console.error);
