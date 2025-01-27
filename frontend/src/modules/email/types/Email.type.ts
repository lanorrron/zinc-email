export type EmailType = {
    message_id: string;
    date: string;
    from: string;
    to: string;
    subject: string;
    mime_version: string;
    content_type: string;
    content_transfer_encoding: string;
    x_from: string;
    x_to: string;
    x_cc: string;
    x_bcc: string;
    x_folder: string;
    x_origin: string;
    x_filename: string;
    cc: string;
    body: string;
  };
  
  export type SearchRequestType = {
    query?: string;
    limit?: number;
    offset?: number;
    startDate?: string;
    endDate?: string;
    nameIndex: string;
  }

  type Hit = {
    _id: string;
    _score: number;
    _source: EmailType;
  }
  
  type Hits = {
    hits: Hit[];
    total: {
      value: number;
    };
    max_score: number;
  }
  
  export type SearchResponseType = {
    hits: Hits;
    timed_out: boolean;
    took: number;
  }