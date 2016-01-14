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
        // We dedicate a method to retrieving and setting some data
        fetchEntries: function() {
            this.entryEndpoint.get().then(function(response) {
                this.$set('entries', response.data.data);
            }, function(response) {
                console.log(response);
            });
        },

        // Adds an event to the existing events array
        postEntry: function() {
            var wrapper = {
                data: this.entry
            }
            this.entryEndpoint.save(wrapper).then(function(response) {
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
