

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
    methods: {
        sendMessage() {
            this.message = ''
        }
    }
})