
const URL = 'https://jsonplaceholder.typicode.com/todos/:afn'


export default function getDominios (name) {
  console.log('name init', name)
  name = name == 'www.google.com' ? 1 : name
  console.log('name end', name)
  const url = URL.replace(':afn', name)
  console.log(URL)
  console.log(url)
  return fetch(URL)
     .then(json => console.log())
}

