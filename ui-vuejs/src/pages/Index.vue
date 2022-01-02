<template>
  <chat/>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import Chat from 'components/chat/chat.vue';

export default defineComponent({
  name: 'PageIndex',
  components: {Chat},
  setup() {
    return {};
  }
});
</script>
