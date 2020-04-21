import _ from 'lodash';
import amqp from 'amqplib'

export function Consumer(msg: amqp.ConsumeMessage, channel: amqp.Channel) {

    console.log("Action received", msg.content.toString());
  
    switch (_.trim(msg.fields.routingKey)) {
      default:
        channel.ack(msg)
        console.log("Event ignored : " + msg.fields.routingKey)
        break;
    }
  }

  export const EventHandlers = {
    "Hello" : function(msg: amqp.ConsumeMessage, channel: amqp.Channel) {

    }
  }