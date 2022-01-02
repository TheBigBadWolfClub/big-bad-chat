import {ChatRoomType, RoomType} from 'src/store/ChatServer/model';

export function AvatarTools() {
  const getImgUrl = (model: ChatRoomType) => {
    if (model.type === RoomType.PUBLIC)
      return `https://robohash.org/${model.name}`

    if (model.type === RoomType.MEETING)
      return `https://robohash.org/${model.name}?set=set4`

    return `https://i.pravatar.cc/150?u=${model.uuid}`;
  }

  return {
    getImgUrl
  }
}
