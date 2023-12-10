<template>
  <a href="/profile">
    <img
      :src="user.image"
      alt="User image"
      style="width: 45px; height: 45px; border-radius: 50%"
  /></a>
</template>

<script>
export default {
  data() {
    return {
      user: {
        // This is a placeholder image. Replace this with the user's image from the backend.
        image: "https://i.pravatar.cc/300",
      },
    };
  },
};
</script>
