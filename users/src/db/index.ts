import mongoose from 'mongoose';
import defaultConfig from '../config.json'

const connectionString= process.env.MONGOURL || defaultConfig.mongoURL;

mongoose.connect(connectionString, {useNewUrlParser: true, useUnifiedTopology: true}).then(()=> {
    console.log("Successfully connected to " + connectionString);
}).catch((err) => {
    console.log("Failed to connect", err)
})