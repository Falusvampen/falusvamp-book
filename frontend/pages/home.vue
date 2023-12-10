<template>
  <div class="homepage">
    <h1>Welcome to My Social Network</h1>
    <p>Connect with friends and share your updates.</p>
    <button @click="login">Log In</button>
    <button @click="signup">Sign Up</button>
  </div>
</template>

<script>
export default {
  methods: {
    login() {
      // Implement your login logic here
    },
    signup() {
      // Implement your signup logic here
    },
  },
};
</script>

<style scoped>
.homepage {
  text-align: center;
  margin-top: 100px;
}

h1 {
  font-size: 32px;
  margin-bottom: 20px;
}

p {
  font-size: 18px;
  margin-bottom: 30px;
}

button {
  padding: 10px 20px;
  font-size: 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  margin-right: 10px;
  cursor: pointer;
}
</style>
