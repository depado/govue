new Vue({
    el: '#entries',
    data: {
        entry: {
            type: 'entry',
            attributes: {
                title: '',
                markdown: '',
            }
        },
        entries: [],
    },
    ready: function() {
        this.entryEndpoint = this.$resource('api/v1/entry/{id}')
        this.fetchEntries();
    },
    methods: {
        fetchEntries: function() {
            this.entryEndpoint.get().then(function(response) {
                this.$set('entries', response.data.data);
            }, function(response) {
                console.log(response);
            });
        },
        postEntry: function() {
            this.entryEndpoint.save({data: this.entry}).then(function(response) {
                this.entries.push(response.data.data);
            }, function(response) {
                console.log(response);
            });
            this.entry.type.attributes = {
                title: '',
                markdown: '',
            };
        },
        deleteEntry: function(index) {
            this.entryEndpoint.delete({id: this.entries[index].attributes.id}).then(function(response) {
                this.entries.splice(index, 1);
            }, function(response) {
                console.log(response)
            });
        }
    }
});
