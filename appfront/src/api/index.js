
const URL = 'https://jsonplaceholder.typicode.com/todos/1'


export default function getDominios (name) {
  const url = URL.replace(':afn', name)
  console.log(URL)
  console.log(url)
  return fetch(URL)
     .then(json => console.log())
}

