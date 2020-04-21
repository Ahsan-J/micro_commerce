import { getCartChannel } from './index';
import defaultConfig from '../../config.json'

export async function sendHelloWorld() {
    const queue = process.env.QUEUE_CART || defaultConfig.queue.cart;
    const channel = await getCartChannel();
    const msg = "Hello world"
    channel.sendToQueue(queue, Buffer.from(msg));
}