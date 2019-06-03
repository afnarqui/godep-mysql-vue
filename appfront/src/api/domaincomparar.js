const observableDiff = require('deep-diff').observableDiff;

export default function buscardomaincomparar (name) {
  return fetch(name)
     .then(response => response.json())
     .then(json => {

      let data = json
      var valorInicial = []
      var valorfinal = []

      for(var i = 0; i< data.length;i++){
        var host = data[i]["host"]
        var port = data[i]["port"]
        var status = data[i]["status"]
        var hostold = data[i]["hostold"]
        var portold = data[i]["portold"]
        var statusold = data[i]["statusold"]
        console.log(host)
        console.log(portold)
        valorInicial.push({
          "host": host,
          "port": JSON.stringify(port),
          "status": status
       })
       valorfinal.push({
        "host": hostold,
        "port": JSON.stringify(portold),
        "status":statusold
     })

        if(i + 1 === data.length){
          var remoto =valorfinal
          var local = valorInicial
          return new Promise( (resolve, reject) => {
            try {

                var remotoA = JSON.parse(JSON.stringify(remoto));
                observableDiff(local, remoto, function (d) {
                    if(d['path'] && d['path'].join('.').indexOf('._id') > -1){
                    }else{

                    }

                    if(d['kind'] == 'A'){

                        if(d['item']['kind'] == 'N'){
                            // new items
                            d['item']['rhs']['_kind_NN'] = 'N';

                            remotoA[d['index']] = "Valor Agregado: " + d['item']['rhs'];
                        }else if(d['item']['kind'] == 'D'){
                            // delete items
                          if(remotoA[d['index']]){
                                remotoA[d['index']]['_kind_DD'] = 'D'
                            }else {
                                remotoA.push(d['item']['lhs']);
                                remotoA[d['index']]['_kind_DD'] = 'D'
                                return;
                            }

                        }
                        return;
                    }

                    var ruta = d['path'];
                    var path = '';
                    var pathInicial = '';
                    var nameInicial = '';

                    let valorEditar = ruta[0]
                    let valorCadena = ruta[1]
                    for (var a = 0; a < ruta.length; a++) {
                        var element = ruta[a];

                        if(a == 0){
                            pathInicial = element;
                        }else if( a == 1){
                            nameInicial = element;
                        }

                        path += '["'+element+'"]';
                    }
                    if(d['kind'] == 'E'){
                        // edit items

                        remotoA[valorEditar][valorCadena] = "Valor Editado: " + d['rhs']

                        // console.log('remotoA'+path+' = `' + d['rhs'] + '`');
                        // console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "E"');

                    }else if(d['kind'] == 'N'){
                        // new columns
                        // console.log('remotoA'+path+' = "' + d['rhs'] + '"');
                        // console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "N"');

                    }else if(d['kind'] == 'D'){
                        // delete;
                        //console.log('remotoA['+pathInicial+']["_kind_'+nameInicial+'"] = "D"');
                    }
                });
                resolve([local, remotoA]);
            } catch (error) {
                console.dir(error)
                reject(error)
            }
        })
        }
      }
         return json
      }
    )
}
