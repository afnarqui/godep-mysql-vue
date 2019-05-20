import { ruta } from './config';

const { api } = ruta

const URL = `${api}`

export default function getDominios (name) {
  const url = URL.replace(':afn', name)
  return fetch(URL)
     .then(json => {
      console.log(json)
      return json.endpoints
     } )
}

