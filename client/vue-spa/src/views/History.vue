<template>
  <div>
    <h1>Yamabiko History</h1>
    <p v-for="(history,index) in yamabiko.history" :key="index+history.textt">
      {{index}} : {{history.text}}
    </p>
  </div>
</template>

<script>
export default {
  name: 'Yamabiko',
  data() {
    return {
      message: '',
      yamabiko: {
        history: [],
      },
    };
  },
  created() {
    this.getHistory();
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
    async getHistory() {
      const url = 'http://localhost:8000/yamabiko';
      try {
        this.yamabiko = await window.fetch(url, {
          method: 'GET',
          headers: {
            'X-Requested-With': 'csrf', // csrf header
            'Content-Type': 'application/json',
          },
        }).then((res) => res.json());
      } catch (e) {
        console.log(e);
      }
    },
  },
};
</script>
