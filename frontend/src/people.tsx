import { createEffect, createResource, For, Show } from "solid-js";
import { backend_url } from "./settings";
import { createMutable } from "solid-js/store";

// const backend_url = "http://localhost:8000"

export interface Person_t {
    id: number;
    name: string;
    parent_id: number;
    image: string;
    gender: string;
    is_descendant: boolean;
    spouse_id: number;
}


// export const [people, {mutate, refetch}] = createResource(() => fetch(`${backend_url}/people`).then(res => res.json()))
export const people = createMutable({} as {[key: number]: Person_t})
window.people = people


fetch(`${backend_url}/people`).then(res => res.json()).then(data => {
    for (const [id, person] of Object.entries(data)) {
        people[Number(id)] = person as Person_t
    }
})
function getChildren(person: Person_t) {
    return Object.values(people).filter(p => p.parent_id === person.id);
}


export const PersonComponent = (props: {person: Person_t}) => {
    const children = () => getChildren(props.person)
    return (
        <div style={{display: "flex", "flex-direction": "column", "align-items": "center"}}>
            <div class="profile" style={{display: "flex", "align-items": "center", gap: "10px", border: "solid 1px black", "border-radius": "5px", padding: "10px"}}>
                <img src={props.person.image} alt="" style={{width: "50px", height: "50px", "border-radius": "50%"}} />
                {props.person.name}
            </div>
                <Show when={children().length} keyed >
            <div style={{display: "flex", gap: "10px", "justify-content": "space-around", border: "solid 1px black", "border-radius": "5px", padding: "10px"}} class="children">
                    <For each={children()}>{(child) => (
                        <PersonComponent person={child} />
                    )}</For>
            </div>
                </Show>
        </div>
    );
}


createEffect(function(){
    console.table(people)
})

export default function People(){
    return(
        <Show when={people[81]} fallback={<div>Loading...</div>}>
            <PersonComponent person={people[81]}  />
        </Show>
    )
}