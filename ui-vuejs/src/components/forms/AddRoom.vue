<template>

  <q-input
    v-model="roomName"
    :rules="[ val => val && val.length > 0 ]"
    bg-color="grey-3"
    bottom-slots
    color="grey-10"
    dense
    label="Add Room"
    label-color="black"
    lazy-rules
    outlined>


    <template v-slot:after>
      <q-btn
        :icon="lockIcon"
        dense
        outline
        rounded
        square
        @click="setPublic"
      />

      <q-btn
        color="secondary"
        dense
        icon="add"
        square
        type="submit"
        @click="onSubmit"
      />
    </template>
  </q-input>


</template>

<script lang="ts">
import {computed, defineComponent, ref} from 'vue';
import {useQuasar} from 'quasar';
import {RoomType} from 'src/store/ChatServer/model';
import {api} from 'boot/axios';


export default defineComponent({
  name: 'AddRoom',
  components: {},
  setup() {
    const $q = useQuasar();

    const onSubmit = (() => {
      const room = {
        name: roomName.value,
        type: isPublic.value ? RoomType.PUBLIC : RoomType.MEETING
      }

      api.post('/rooms', room)
        .then(() => $q.notify('Room  Created'))
        .catch((err) => $q.notify(`Fail to create room : ${err}`))
    })

    const setPublic = (() => {
      isPublic.value = !isPublic.value
    })
    const roomName = ref('')
    const isPublic = ref(false)
    const lockIcon = computed(() => {
      return isPublic.value ? 'lock_open' : 'lock'
    })
    return {onSubmit, roomName, isPublic, setPublic, lockIcon}
  }
})
</script>



