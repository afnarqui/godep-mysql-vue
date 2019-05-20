
import axios from "axios";

new Vue({
    el: '#App',
    created() {

    },
    data: {
        message: '',
        messages: [{
            text: 'hello',
            date: new Date()
        }]
    },
    mounted() {
            axios({ method: "GET", "url": "http://localhost:8081/notes" }).then(result => {
                console.log(result.data.origin)    
                this.messages = result.data.origin;
            }, error => {
                console.error(error);
            });
        },    
    methods: {
        sendMessage() {
            this.message = ''
        }
    }
})