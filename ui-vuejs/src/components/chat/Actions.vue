<template>
  <div>
    <q-fab
      v-model="fab2"
      color="brown"
      direction="down"
      glossy
      icon="more_vert"
      padding="xs"
      vertical-actions-align="right"
    >
      <q-fab-action color="primary" icon="mail" label="Email" label-position="left" @click="onClick"/>
      <q-fab-action color="secondary" icon="alarm" label="Alarm" label-position="left" @click="onClick"/>
      <q-fab-action color="orange" icon="airplay" label="Airplay" label-position="left" @click="onClick"/>
      <q-fab-action color="accent" icon="room" label="Map" label-position="left" @click="onClick"/>
    </q-fab>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue';
import {logger} from 'boot/logger';

export default defineComponent({
  name: 'RoomActions',
  setup() {
    const fab2 = ref(false)

    const onClick = () => logger.warn('missing: run action')
    return {fab2, onClick}
  },
})
</script>

<style lang="sass" scoped>

</style>

