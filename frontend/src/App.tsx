import { createResource, For, type Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';
import { createMutable } from 'solid-js/store';
import People from './people';



// const backend_url = "http://localhost:8000"
const backend_url = ""
const ws = new WebSocket(`${backend_url}/ws`)


const [todos, {mutate}] = createResource(() => fetch(`${backend_url}/todos`).then(res => res.json()))
const messages = createMutable([])


ws.onmessage = (e) => {
  messages.push(e.data)
}



  const App: Component = () => {
  return (
    <div class={styles.App}>
      <For each={todos()}>{(todo) => (
        <div>
          <p>{todo.title}</p>
        </div>
      )}</For>
      <For each={messages}>{(message) => (
        <div>
          <p>{message}</p>
        </div>
      )}</For>
      <form action="" onSubmit={e => {
        e.preventDefault()
        const message = (e.target as HTMLFormElement).message.value
        ws.send(message)
      }}>
        <input type="text" name='message' />
        <button type="submit">Send</button>
      </form>
      <People/>
    </div>
  );
};

export default App;
