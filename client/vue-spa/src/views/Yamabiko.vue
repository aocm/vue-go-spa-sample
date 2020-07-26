<template>
  <div>
    <h1>This is an Yamabiko page</h1>
    <input v-model="message" placeholder="Say Yahho">
    <button @click="Send">Send</button>
    <p>Yamabiko : {{ yamabiko }}</p>
  </div>
</template>

<script>
export default {
  name: 'Yamabiko',
  data() {
    return {
      message: '',
      yamabiko: '',
    };
  },
  methods: {
    async Send() {
      const yamabiko = await this.CallYamabikoAPI().then((res) => res.json());
      this.yamabiko = yamabiko.message;
      window.alert(this.yamabiko);
    },

    /**
     * TODO 共通化させる
     */
    async CallYamabikoAPI() {
      const url = 'http://localhost:8000/yamabiko';
      const data = {
        message: this.message,
      };
      try {
        return await window.fetch(url, {
          method: 'POST',
          headers: {
            'X-Requested-With': 'csrf', // csrf header
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });
      } catch (e) {
        console.log(e);
        return e;
      }
    },
  },
};
</script>
