import { createResource, For, type Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';
import { createMutable } from 'solid-js/store';
import People, { people, Person_t } from './people';
import { backend_url } from './settings';



// const backend_url = "http://localhost:8000"

const ws = new WebSocket(`${backend_url}/ws`)




const [todos, {mutate}] = createResource(() => fetch(`${backend_url}/todos`).then(res => res.json()))
const messages = createMutable([])


ws.onmessage = (e) => {
  console.log("data", e.data)
  const json = JSON.parse(e.data)
  const message_handlers: {[key: string]: (json: any) => void} = {
    row_added,
    row_updated,
    store_join
  }
  const handler = message_handlers[json.type]
  if (handler) {
    handler(json)
  }
}



//start
function row_added(json: {row: any, id: number}){
  people[json.id as number] = json.row
}

function row_updated(json: {row: any, id: number}){
  people[json.id as number] = json.row
}

function store_join(json: {rows: any}){
  for (const [id, obj] of Object.entries(json.rows as {[key: number]: Person_t})) {
    people[Number(id)] = obj
  }
}
//end

  const App: Component = () => {
  return (
    <div class={styles.App}>
      <People/>
    </div>
  );
};

export default App;
