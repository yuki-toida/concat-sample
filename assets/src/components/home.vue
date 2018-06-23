<template>
  <div>
    <div>
      <label>
        <input type="file" id="files" ref="files" accept="video/*" multiple v-on:change="change()">
      </label>
    </div>
    <div>
      <div v-for="(file, key) in files" v-bind:key="key">
        {{ file.name }}
      </div>
    </div>
    <div>
      <button v-on:click="upload()">Upload</button>
    </div>
    <div v-if="output">
      <a v-bind:href="link" download="output.mp4">download</a>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data: function() {
    return {
      domain: "http://localhost:8080",
      files: [],
      output: null
    }
  },
  computed: {
    link: function() {
      return `${this.domain}/static/video/${this.output}`;
    }
  },
  methods: {
    change: function() {
      this.files = this.$refs.files.files;
    },
    upload: function() {
      let formData = new FormData();
      for( var i = 0; i < this.files.length; i++ ) {
        formData.append(`file${i}`, this.files[i]);
      }

      axios.post(`${this.domain}/upload`, formData, {headers: {'Content-Type': 'multipart/form-data'}})
      .then((res) => this.output = res.data)
      .catch((error) => alert(error));
    }
  }
}
</script>
