import { createEffect, createResource, For, Show } from "solid-js";

// const backend_url = "http://localhost:8000"
const backend_url = ""

export interface Person {
    id: number;
    name: string;
    parent_id: number;
}


const [people, {mutate, refetch}] = createResource(() => fetch(`${backend_url}/people`).then(res => res.json()))

function getChildren(person: Person) {
    return people().filter(p => p.parent_id === person.id);
}


export const PersonComponent = (props: {person: Person}) => {
    return (
        <div>
            {props.person.name}
            <div style={{
                display: "flex",
                gap: "10px",
                "justify-content": "space-around",
                "align-items": "center",
                border: "solid 1px black",
                padding: "10px",
                "border-radius": "5px"
            }} class="children">
                <For each={getChildren(props.person)}>{(child) => (
                    <PersonComponent person={child} />
                )}</For>
            </div>
        </div>
    );
}


createEffect(function(){
    console.table(people())
})

export default function People(){
    return(
        <Show when={people()} fallback={<div>Loading...</div>}>
            <PersonComponent person={people().find(p => p.id == 1)}  />
        </Show>
    )
}