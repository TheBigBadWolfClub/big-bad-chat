<template>

  <q-form
    class="column serverForm"
    @submit="onSubmit"
  >
    <q-chip class="label" square>
      <q-avatar color="red" icon="bookmark" text-color="white"/>
      Server Config
    </q-chip>


    <div class="row content-stretch formRow">
      <q-input
        v-model="srvConf.addr"
        :rules="[ val => val && val.length > 0 ]"
        bg-color="grey-3"
        class="col"
        color="grey-10"
        dense
        label="Server Addr"
        label-color="black"
        lazy-rules
        outlined
      />
      <div class="spacer"></div>
      <q-input
        v-model="srvConf.port"
        bg-color="grey-3"
        class="col-4"
        color="grey-10"
        dense
        label="Server Port"
        label-color="black"
        lazy-rules
        outlined
      />
    </div>

    <div class="formRow">
      <q-item-label
        class="text-grey-8"
        header
      >Rest
      </q-item-label>

      <div class="row content-stretch no-wrap formRow">
        <q-input
          v-model="srvConf.rest.protocol"
          :rules="[ val => val && val.length > 0 ]" bg-color="grey-3"
          color="grey-10"
          dense
          hint="http"
          label="Protocol"
          label-color="black"
          lazy-rules
          outlined
        />
        <div class="spacer"></div>
        <q-input

          v-model="srvConf.rest.uri"
          :rules="[ val => val && val.length > 0 ]" bg-color="grey-3"
          color="grey-10"
          dense
          hint="/api"
          label="Uri"
          label-color="black"
          lazy-rules
          outlined
        />
      </div>
    </div>

    <div class="formRow">
      <q-item-label
        class="text-grey-8"
        header
      >Socket
      </q-item-label>
      <div class="row content-stretch no-wrap formRow">

        <q-input
          v-model="srvConf.webSocket.protocol"
          :rules="[ val => val && val.length > 0 ]" bg-color="grey-3"
          color="grey-10"
          dense
          hint="ws"
          label="Protocol"
          label-color="black"
          lazy-rules
          outlined
        />
        <div class="spacer"></div>
        <q-input

          v-model="srvConf.webSocket.uri"
          :rules="[ val => val && val.length > 0 ]" bg-color="grey-3"
          color="grey-10"
          dense
          hint="/ws"
          label="Uri"
          label-color="black"
          lazy-rules
          outlined
        />
      </div>
    </div>

    <div class="row justify-end formRow">
      <q-btn class="submitBtn  col-4"
             color="secondary"
             label="Update"
             padding="none"
             text-color="white"
             type="submit"
      />
    </div>


  </q-form>
</template>

<script lang="ts">
import {useStore} from 'src/store';
import {computed, defineComponent, WritableComputedRef} from 'vue';
import {useQuasar} from 'quasar';
import {AppConfig, AppConfigId, ServerConfig} from 'src/store/AppConfig/model';
import {SocketConfig, SocketConfigId} from 'src/store/SocketConfig/model';


export default defineComponent({
  name: 'ConfigForm',
  components: {},
  setup() {
    const store = useStore();
    const $q = useQuasar();


    const srvConf: WritableComputedRef<ServerConfig> = computed(({
      get: () => {
        return store.getters[AppConfigId(AppConfig.Getter.ServerConfig)]
      },
      set: (v: ServerConfig) => {
        void store.dispatch(AppConfigId(AppConfig.Mutation.SetServerConfig), v)
      }
    }))

    const onSubmit = (async () => {
      await store.dispatch(SocketConfigId(SocketConfig.Action.Disconnect))
      $q.notify('Server config updated, Reconnect server.')
    })


    return {onSubmit, srvConf}
  }
})
</script>


<style lang="sass" scoped>

.formRow
  padding-left: 10px
  padding-right: 10px


.spacer
  width: 10px

.submitBtn
  margin-top: 10px

</style>

