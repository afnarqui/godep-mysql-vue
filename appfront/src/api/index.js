import { ruta } from './config';

const { api } = ruta

const URL = `${api}`

export default function getDominios (name) {
  const url = URL.replace(':afn', name)
  console.log(URL)
  console.log(url)
  return fetch(URL)
     .then(json => console.log())
}

