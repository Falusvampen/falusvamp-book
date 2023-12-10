<template>
  <div class="friend-profile">
    <img :src="friend.picture" :alt="friend.name" class="friend-picture" />
    <p class="friend-name">{{ friend.name }}</p>
  </div>
</template>

<script>
export default {
  props: {
    friend: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style scoped>
.friend-profile {
  display: flex;
  align-items: center;
}

.friend-picture {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-right: 10px;
}

.friend-name {
  font-weight: bold;
}
</style>
