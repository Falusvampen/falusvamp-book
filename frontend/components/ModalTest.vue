<template>
  <div class="modal-wrapper" @click="$emit('close')">
    <div class="modal-container" :class="size" @click.stop="">
      <button @click="$emit('close')">Close</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    size: {
      type: String,
      default: "",
    },
  },
};
</script>
