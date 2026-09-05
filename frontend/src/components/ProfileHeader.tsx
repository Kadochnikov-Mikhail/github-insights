type ProfileHeaderProps = {
  username: string;
};

function ProfileHeader({ username }: ProfileHeaderProps) {
  return (
    <div className="profile-header">
      <div className="profile-info">
        <div className="profile-avatar">
          {username.charAt(0).toUpperCase()}
        </div>

        <div>
          <h2>{username}</h2>
          <p>GitHub profile analysis</p>
        </div>
      </div>

      <div className="profile-status">
        <span />
        Synced just now
      </div>
    </div>
  );
}

export default ProfileHeader;