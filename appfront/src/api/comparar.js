// import { uuid } from 'vue-uuid';
// const observableDiff = require('deep-diff').observableDiff;
// export const Uuid = uuid.v1();
// console.log('UUid :' + Uuid)
// console.log(Uuid);

// let dato = [{ 'nombre': 'andres', 'quehago': false, 'rutas': 'c:\\bd', 'tabla': 'prueba' }]
// let dato2 = [{ 'apellido': 'naranjo', 'edad': 28, 'sexo': 'F', 'quehago': true, 'nombre': 'andres', 'rutas': 'c:\\bd' }]

//  function comparar(local, remoto) {

//   return new Promise( (resolve, reject) => {
//       try {
//           console.log('entro a promise')
//           console.log('remot',remoto)
//           console.log('local',local)
//           var remotoA = JSON.parse(JSON.stringify(remoto));
//           observableDiff(local, remoto, function (d) {
//               console.log('datos d:')
//               console.log(d)
//               if(d['path'] && d['path'].join('.').indexOf('._id') > -1){
//               }else{

//               }

//               if(d['kind'] == 'A'){

//                   if(d['item']['kind'] == 'N'){
//                       // indicar que la accion es un nuevo registro
//                       d['item']['rhs']['_kind_NN'] = 'N';
//                       // remotoA.push(d['item']['rhs']);
//                       remotoA[d['index']] = d['item']['rhs'];
//                   }else if(d['item']['kind'] == 'D'){
//                       // indicar que la accion es eliminar
//                       // saber si se elimino del remoto o del local
//                       if(remotoA[d['index']]){
//                           remotoA[d['index']]['_kind_DD'] = 'D'
//                       }else {
//                           remotoA.push(d['item']['lhs']);
//                           remotoA[d['index']]['_kind_DD'] = 'D'
//                           return;
//                       }

//                   }
//                   return;
//               }

//               var ruta = d['path'];
//               var path = '';
//               var pathInicial = '';
//               var nameInicial = '';
//               // encontar objecto a editar
//               console.log('que tiene ruta:')
//               console.log(ruta)
//               for (var a = 0; a < ruta.length; a++) {
//                   var element = ruta[a];

//                   if(a == 0){
//                       pathInicial = element;
//                   }else if( a == 1){
//                       nameInicial = element;
//                   }

//                   path += '["'+element+'"]';
//               }
//               console.log('resultado final path')
//               console.log(path)
//               if(d['kind'] == 'E'){
//                   // editar

//                   console.log('remotoA'+path+' = `' + d['rhs'] + '`');
//                   // console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "E"');
//                   console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "E"');

//               }else if(d['kind'] == 'N'){
//                   // crear nuevo valor de columna
//                   console.log('remotoA'+path+' = "' + d['rhs'] + '"');
//                   console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "N"');

//               }else if(d['kind'] == 'D'){
//                   // eliminar;
//                   console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "D"');
//               }
//           });
//           console.log('resultado final')
//           console.log(local,remotoA);
//           resolve([local, remotoA]);
//       } catch (error) {
//           console.dir(error)
//           reject(error)
//       }
//   });
// }

export default function armarCadena(data) {
  debugger
  let valorInicial = []
  let valorfinal = []
  var promise = new Promise((resolve, reject ) =>{

    console.log(data)
  if(data.length==0){
    return reject([],[])
  }
  let valorInicial = []
  let valorfinal = []
  for(var i = 0; i< data.length;i++){
    valorInicial.push({
      host: data[i]['host'],
      port: data[i]['port'],
      status: data[i]['status']
    })
    valorFinal.push({
      host: data[i]['hostold'],
      port: data[i]['portold'],
      status: data[i]['statusold']
    })
    if(i === data.length + 1){
      console.log(valorInicial,valorfinal)
    }
  }
  debugger
  console.log(valorInicial)
  console.log(valorFinal)


  })

  return promise

}
