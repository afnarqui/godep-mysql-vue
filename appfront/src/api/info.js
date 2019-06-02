
const URL = 'https://jsonplaceholder.typicode.com/todos/:afn'


export default function getDominios (name) {
  // console.log('name end', name)
  const url = URL.replace(':afn', name)
  // console.log(name)
  // console.log(url)
  debugger
  return fetch(name)
     .then(response => response.json())
     .then(json => {
       return json
      }
       )
}
