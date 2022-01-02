<template>
  <div class="q-pa-md" style="max-width: 350px">
    <q-list>
      <CategoryLayout v-for="category in groupByCategory" :key="category" v-bind:model=category.layoutOptions>
        <ChatEntryLayout v-for="chatRoom in category.dataSource.value"
                         :key="chatRoom"
                         v-bind:model="chatRoom"/>
      </CategoryLayout>
    </q-list>
  </div>
</template>
<script lang="ts">

import {computed, defineComponent} from 'vue';
import CategoryLayout, {CategoryOptions} from 'components/ChatList/categoryLayout.vue';
import {ChatServer, ChatServerId} from 'src/store/ChatServer/model';
import {useStore} from 'src/store';
import ChatEntryLayout from 'components/ChatList/chatEntryLayout.vue';

interface Category {
  order: number;
  dataSource: any;
  layoutOptions: CategoryOptions
}

export default defineComponent({
  name: 'ChatRooms',
  components: {ChatEntryLayout, CategoryLayout},
  setup() {
    const store = useStore();

    const directChat: Category = {
      order: 1,
      dataSource: computed(() => store.getters[ChatServerId(ChatServer.Getter.DirectRooms)]),
      layoutOptions: {label: 'CoWorkers', icon: 'perm_identity', caption: 'direct message'}
    };

    const meetingChat: Category = {
      order: 2,
      dataSource: computed(() => store.getters[ChatServerId(ChatServer.Getter.MeetingRooms)]),
      layoutOptions: {label: 'Meetings', icon: 'groups', caption: 'available meetings'}
    }

    const groupChat: Category = {
      order: 3,
      dataSource: computed(() => store.getters[ChatServerId(ChatServer.Getter.PublicRooms)]),
      layoutOptions: {label: 'Groups', icon: 'group_add', caption: 'join discussion groups'}
    }

    const groupByCategory: Category[] = []
    groupByCategory.push(directChat, meetingChat, groupChat)
    return {groupByCategory}
  }

})
</script>

<style lang="sass">
.entry
  width: 100%

</style>

