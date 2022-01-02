<template>

  <q-chip class="entry row justify-between " square>
    <q-avatar class="avatar">{{ room.is_public }}</q-avatar>
    <div class="col">{{ room.name }}</div>

    <q-btn :loading="isLoading" class="join-btn"
           color="green"
           icon="check"
           padding="none"
           round
           @click="onJoin(room)">
      <template v-slot:loading>
        <q-spinner-gears/>
      </template>
    </q-btn>

  </q-chip>


</template>

<script lang="ts">
import {defineComponent, PropType, ref} from 'vue';
import {joinRoomMsgString} from 'src/store/SocketConfig/message';
import {useStore} from 'src/store';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';
import {ChatRoomType} from 'src/store/ChatServer/model';


export default defineComponent({
  name: 'RoomEntry',
  props: {
    room: {type: Object as PropType<ChatRoomType>, required: true},
    active: {
      type: Boolean
    }
  },
  setup() {
    const store = useStore();
    const isLoading = ref(false);
    const onJoin = (roomId: string) => {
      isLoading.value = true
      const msgStr = joinRoomMsgString(roomId);
      const socketConn = store.getters[SocketConfigId(SocketConfig.Getter.getSocket)];
      socketConn.send(msgStr)

      // simulate a delay
      setTimeout(() => {
        // we're done, we reset loading state
        isLoading.value = false
      }, 1000)
    }
    return {onJoin, isLoading}
  },
})
</script>

<style lang="sass" scoped>

.chip
  background-color: white
  color: black

.avatar
  background-color: red
  color: white


.join-btn
  background-color: white
  color: black
  margin-right: 0
  float: right
  justify-content: right

</style>

