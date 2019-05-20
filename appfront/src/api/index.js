import { ruta } from './config';

const { api } = ruta

const URL = `${api}`

export default function getDominios (name) {
  const url = URL.replace(':afn', name)
  return fetch(url)
     .then(json => json.endpoints)
}

