
const socket = io();

new Vue({
    el: '#chat-app',
    created() {

    },
    data: {
        message: '',
        messages: [{
            text: 'hello',
            date: new Date()
        }]
    },
    methods: {
        sendMessage() {
            socket.emit('chat', this.message)
            this.message = ''
        }
    }
})