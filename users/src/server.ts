import koa from 'koa';
import mount from 'koa-mount';
import graphqlHTTP from 'koa-graphql';
import {schema} from './schema';
import defaultConfig from './config.json'
import './db'

const app = new koa();
const port = (process.env.PORT || defaultConfig.port) + ""

app.use(mount('/graphql', graphqlHTTP({
  schema,
  graphiql:true
})))


app.listen(parseInt(port), ()=> {
  console.log("Listening at port ", port)
});