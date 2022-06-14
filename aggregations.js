db.podcasts.aggregate(
  [
    { $unwind: "$episodes" },
    { $match: { 'episodes.title': /Holmes/ } },
    {
      $project: {
        "podcast": "$title",
        "episode": "$episodes.title",
        "duration": "$episodes.duration",
        "createdAt": "$episodes.createdAt"
      }
    }
  ]
)
