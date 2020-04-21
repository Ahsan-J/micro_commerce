import amqp from 'amqplib';
import { validate } from '../helper/utility';
import defaultConfig from '../config.json';
import { registerCartQueue } from './cart';

let cacheConnection: amqp.Connection;

export const connectToRabbitMQ = () => {
    const src = process.env.rabbitMQ_URL || defaultConfig.rabbitMQ_URL
    amqp.connect(src).then(async (connection: amqp.Connection) => {
        cacheConnection = connection;
        console.log("Connection Established to RabbitMQ")
        registerCartQueue(await getChannel());
        
    }).catch((connectionError) => {
        console.log("Connection Error connecting to rabbitMQ", connectionError)
    });
}

export const getChannel = (connection: amqp.Connection = cacheConnection) => new Promise<amqp.Channel>((resolve, reject) => {
    if(!validate(connection)) {
        connectToRabbitMQ()
        getChannel();
    }

    connection.createChannel().then((channel: amqp.Channel) => {
        resolve(channel);
    }).catch((channelError) => {
        console.log("Channel error to rabbitMQ", channelError)
        reject(channelError);
    });
})