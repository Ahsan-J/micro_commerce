import amqp from 'amqplib';
import _ from 'lodash';
import defaultConfig from '../../config.json';
import { Consumer } from './receiver';
import { validate } from '../../helper/utility';
import { getChannel } from '../index';

const Events: Array<string> = ['hello']

let cartChannel: amqp.Channel;

export const registerCartQueue = async (channel: amqp.Channel) => {
    
    if(!validate(channel)) {
        channel = await getChannel();
    }

    const exchange = process.env.EXCHANGE_CART || defaultConfig.exchange.cart
    const queue = process.env.QUEUE_CART || defaultConfig.queue.cart

    channel.assertExchange(exchange, 'topic', { durable: true });
    channel.assertQueue(queue, {exclusive : false, durable:true});
    // binding Queues

    _.forEach(Events, (eventName) => {
        channel.bindQueue(queue, exchange, eventName);
    })

    channel.consume(queue, (msg)=> msg ? Consumer(msg,channel) : console.log("Message Empty", msg) , { noAck: false });

    cartChannel = channel;
    return channel;
}

export const getCartChannel: () => Promise<amqp.Channel> = async () => {
    if(!validate(cartChannel)) {
        return registerCartQueue(await getChannel())
    }
    return cartChannel;
}