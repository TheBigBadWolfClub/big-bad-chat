<template>
  <q-card-section class="col bg-grey-9 full-height full-width">
    <div class="q-pa-md row justify-center">
      <div class="full-width">
        <ChatEntry v-for="msg in messages" :key="msg.value.uuid" ref="msg"
                   :message="msg"
                   :room="room"
                   :user="user"
        />
      </div>
    </div>
  </q-card-section>
</template>

<script lang="ts">
import {useStore} from 'src/store';
import {computed, defineComponent, reactive, Ref, ref, watchEffect} from 'vue';
import {ChatMessage, MsgChat, MsgType} from 'src/store/SocketConfig/message';
import ChatEntry from 'components/chat/ChatEntry.vue';
import {ChatServer, ChatServerId} from 'src/store/ChatServer/model';
import {uuid} from 'vue-uuid';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';
import {logger} from 'boot/logger';


export default defineComponent({
  name: 'chat-input',
  components: {ChatEntry},
  setup() {
    const store = useStore();
    const state = reactive({message: ''})
    const received = reactive(['received'])
    const send = reactive(['send'])

    const testMsg: Ref<ChatMessage> = ref({
      uuid: uuid.v4(),
      sender: {
        uuid: '',
        name: '',
        connection_uuid: ''
      },
      room_uuid: '',
      time: 'time',
      text: ['hi', 'john'],
    })

    const testMsg2: Ref<ChatMessage> = ref({
      uuid: uuid.v4(),
      sender: {
        uuid: '21323e',
        name: 'Marta',
        connection_uuid: ''
      },
      room_uuid: '',
      time: 'time',
      text: ['im', 'marta'],
    })

    const user = computed(() => {
      const user = store.getters[ChatServerId(ChatServer.Getter.User)];

      // TODO unk
      //testMsg.value.sender.uuid = user.uuid
      //testMsg.value.sender.name = user.name
      return user
    })

    const room = computed(() => {
      const room = store.getters[ChatServerId(ChatServer.Getter.ActiveChatRoom)];
      // TODO unk
      //testMsg.value.dest_uuid = room.uuid
      //testMsg2.value.dest_uuid = room.uuid
      //testMsg2.value.sender.uuid = room.uuid
      return room
    })


    watchEffect(() => {
      logger.debug("-------------")
      const msg: MsgChat = store.getters[SocketConfigId(SocketConfig.Getter.onMessage)]
      if (msg.action !== MsgType.NewTextAction) return;

      received.push(...msg.chat.text)
    })


    const messages = ref([]);
    return {messages, user, room}
  }
})
</script>

