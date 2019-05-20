
const URL = 'https://api.ssllabs.com/api/v3/analyze?host=www.google.com'


export default function getDominios (name) {
  const url = URL.replace(':afn', name)
  console.log(URL)
  console.log(url)
  return fetch(URL)
     .then(json => console.log())
}

